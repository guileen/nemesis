package modules

import "log"

type ServerModule struct {
	Address   string
	Hosts     []string
	IsDefault bool
	Chain     Module
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
			delete(conf, k)
		case "bind":
			m.Address = v.(Scalar).String()
			delete(conf, k)
		case "host":
			for _, h := range v.(List) {
				m.Hosts = append(m.Hosts, h.(Scalar).String())
			}
			delete(conf, k)
		case "default":
			m.IsDefault = v.(Scalar).GetBool()
			delete(conf, k)
		default:
			log.Println("config>", k, ":", v)
		}
	}
	m.Chain = MakeChain(node.(Map))
	m.Address = regularAddress(m.Address)
}

func (m *ServerModule) Process(req *Req, res *Res) bool {
	// TODO Process Modules
	log.Println("server process")
	m.Chain.Process(req, res)
	return false
}
