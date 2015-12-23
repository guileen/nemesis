package modules

import (
	"log"
	"net/http"
)

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
	log.Println("Static process:", req.GetPath())
	m.fileServer.ServeHTTP(res, req.request)
	if res.statusCode >= 400 && res.statusCode < 500 {
		return false
	}
	return true
}
