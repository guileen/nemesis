package modules

import (
	"log"
	"net/http"
)

type RootModule struct {
	Upstreams     map[string]*UpstreamModule
	Servers       map[string]*ServerModule
	DefaultServer *ServerModule
	Addresses     map[string]bool
}

func NewRootModule(root Node) *RootModule {
	module := &RootModule{}
	module.Init(root)
	return module
}

func (root *RootModule) Init(node Node) {
	root.Upstreams = make(map[string]*UpstreamModule)
	root.Servers = make(map[string]*ServerModule)
	root.Addresses = make(map[string]bool)
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
			server := NewServerModule(k, v)
			root.Addresses[server.Address] = true
			root.Servers[k] = server
		}
	}
}

func (root *RootModule) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
}

func (root *RootModule) Run() {
	log.Println(root.Addresses)
	done := make(chan bool)
	for addr := range root.Addresses {
		log.Println("Listen at", addr)
		go http.ListenAndServe(addr, root)
	}
	<-done
}
