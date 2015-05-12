package modules

import (
	"log"
	"regexp"
)

type RoutesModule struct {
	Routes []*Route
}

func NewRoutesModule(node Node) *RoutesModule {
	module := &RoutesModule{
		Routes: make([]*Route, 0),
	}
	module.Init(node)
	return module
}

func (m *RoutesModule) Init(node Node) {
	list := node.(List)
	for i := range list {
		m.Routes = append(m.Routes, NewRoute(list[i]))
	}
}

func (m *RoutesModule) Process(req *Req, res *Res) bool {
	path := req.GetPath()
	log.Println("path", path)
	for _, route := range m.Routes {
		if route.Match(req) {
			route.Process(req, res)
			break
		}
	}
	return false
}

type Route struct {
	Method    string
	Path      string
	PathRegex string
	Modules   []Module
}

func NewRoute(node Node) *Route {
	route := &Route{}
	route.Init(node)
	return route
}

func (r *Route) Init(node Node) {
	mp := node.(Map)
	if len(mp) > 1 || len(mp) == 0 {
		panic("Invalid route config, only allow ONE path for each route.")
	}
	key := mp.Keys()[0]
	value := mp[key]
	r.Modules = MakeModules(value)
	pairs := regexp.MustCompile("\\s+").Split(key)
	r.Method = pairs[0]
	r.Path = pairs[1]
	// TODO ....
	// r.PathRegex = regexp.MustCompile( r.Path )
}

func (r *Route) GetPath() {
}

func (r *Route) Match(req *Req) bool {
	return false
}

func (r *Route) Process(req *Req, res *Res) bool {
	return false
}
