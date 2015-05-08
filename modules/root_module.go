package modules

import "log"

type RootModule struct {
	Upstreams     map[string]*UpstreamModule
	Servers       map[string]*ServerModule
	DefaultServer *ServerModule
}

func NewRootModule(root Node) *RootModule {
	module := &RootModule{}
	module.Init(root)
	return module
}

func (root *RootModule) Init(node Node) {
	root.Upstreams = make(map[string]*UpstreamModule)
	root.Servers = make(map[string]*ServerModule)
	conf := node.(Map)
	for k, v := range conf {
		switch k {
		case "upstream":
			mp, ok := v.(Map)
			if !ok {
				log.Println("upstream config", v)
				panic("upstream config must be Map")
			}
			for kk, vv := range mp {
				root.Upstreams[kk] = NewUpstreamModule(vv)
			}
		case "plugins":
			panic("Not support [" + k + "] yet.")
		default:
			root.Servers[k] = NewServerModule(k, v)
		}
	}
}
