package modules

import "log"

type ServerModule struct {
	Address   string
	Hosts     []string
	IsDefault bool
	Routes    *RoutesModule
	Modules   []Module
}

func NewServerModule(host string, node Node) *ServerModule {
	module := &ServerModule{
		Hosts: make([]string, 0),
	}
	module.Hosts = append(module.Hosts, host)
	module.Init(node)
	return module
}

func regularAddress(addr string) string {
	if addr == "" {
		return "0.0.0.0:80"
	}
	if addr[0:1] == ":" {
		return "0.0.0.0" + addr
	}
	return addr
}

// Init http modules
func (m *ServerModule) Init(node Node) {
	log.Println("ServerModule.Init")
	conf := node.(Map)
	// attach to existing
	for k, v := range conf {
		switch k {
		case "port":
			m.Address = ":" + v.(Scalar).String()
		case "bind":
			m.Address = v.(Scalar).String()
		case "host":
			for _, h := range v.(List) {
				m.Hosts = append(m.Hosts, h.(Scalar).String())
			}
		case "default":
			m.IsDefault = v.(Scalar).GetBool()
		case "routes":
			m.Routes = NewRoutesModule(v)
		default:
			log.Println("config>", k, ":", v)
		}
	}
	m.Address = regularAddress(m.Address)
}

func (m *ServerModule) Process(req *Req, res *Res) bool {
	// TODO Process Modules
	m.Routes.Process(req, res)
	return false
}