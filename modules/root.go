package modules

import (
	"log"
	"net/http"
	"strings"
)

type Listener struct {
	Address       string
	Servers       map[string]*ServerModule
	DefaultServer *ServerModule
}

func NewListener(addr string) *Listener {
	return &Listener{Address: addr, Servers: make(map[string]*ServerModule)}
}

func (lsn *Listener) AddServer(server *ServerModule) {
	if server.IsDefault {
		if lsn.DefaultServer != nil {
			log.Println("current default server ", lsn.DefaultServer)
			panic("Duplicate default server on " + lsn.Address)
		}
		lsn.DefaultServer = server
	}
	for _, host := range server.Hosts {
		if lsn.Servers[host] != nil {
			panic("Duplicate server definition of [" + host + "] at [" + lsn.Address + "]")
		}
		lsn.Servers[host] = server
	}
}

func (lsn *Listener) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("req", request)
	host := request.Host
	idx := strings.LastIndex(host, ":")
	if idx >= 0 {
		host = host[:idx]
	}
	server := lsn.Servers[host]
	if server == nil {
		server = lsn.DefaultServer
	}
	if server == nil {
		log.Println("No match host", request.Host)
	}
	req := &Req{request: request}
	res := &Res{Req: req, writer: &writer}
	req.Res = res
	// Do server module
	server.Process(req, res)
}

func (lsn *Listener) Run() {
	log.Println("Listen at", lsn.Address)
	err := http.ListenAndServe(lsn.Address, lsn)
	if err != nil {
		log.Fatal("Error to listen at "+lsn.Address, err)
	}
}

type RootModule struct {
	Upstreams map[string]*UpstreamModule
	Listeners map[string]*Listener
}

func NewRootModule(root Node) *RootModule {
	module := &RootModule{}
	module.Init(root)
	return module
}

func (root *RootModule) Init(node Node) {
	root.Listeners = make(map[string]*Listener)
	root.Upstreams = make(map[string]*UpstreamModule)
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
			root.getOrMakeListener(server.Address).AddServer(server)
		}
	}
}

func (root *RootModule) getOrMakeListener(addr string) *Listener {
	lsn := root.Listeners[addr]
	if lsn == nil {
		lsn = NewListener(addr)
		root.Listeners[addr] = lsn
	}
	return lsn
}

func (root *RootModule) Run() {
	for _, lsn := range root.Listeners {
		go lsn.Run()
	}
}
