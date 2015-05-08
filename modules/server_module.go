package modules

import "log"

type ServerModule struct {
	Address string
	Hosts   []string
	Routes  RoutesModule
}

func NewServerModule(host string, node Node) *ServerModule {
	module := &ServerModule{
		Hosts: make([]string, 0),
	}
	module.Hosts = append(module.Hosts, host)
	return module
}

// Init http modules
func (m *ServerModule) Init(node Node) {
	conf := node.(Map)
	// attach to existing
	for k, v := range conf {
		switch k {
		case "port":
			m.Address = ":" + v.(string)
		case "address":
			m.Address = v.(string)
		case "host":
			for _, h := range v.(List) {
				m.Hosts = append(m.Hosts, h.(string))
			}
		default:
			log.Println("config[", k, v)
		}
	}
}
