package modules

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

type Route struct {
}

func NewRoute(node Node) *Route {
	return nil
}
