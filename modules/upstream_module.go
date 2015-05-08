package modules

import (
	"log"
	"net/http/httputil"
	"net/url"
)

type UpstreamModule struct {
	rproxys []*httputil.ReverseProxy
}

func NewUpstreamModule(node Node) *UpstreamModule {
	module := &UpstreamModule{}
	module.Init(node)
	return module
}

func (up *UpstreamModule) AddNode(node Node) {
	str, _ := node.(string)
	upurl, err := url.Parse(str)
	if err != nil {
		panic(err.Error())
	}
	up.rproxys = append(up.rproxys, httputil.NewSingleHostReverseProxy(upurl))
}

func (up *UpstreamModule) Init(node Node) {
	up.rproxys = make([]*httputil.ReverseProxy, 0)
	str, ok := node.(Scalar)
	if ok {
		up.AddNode(str)
	} else {
		list, ok := node.(List)
		if !ok {
			log.Println("Upstream config", node)
			panic("Upstream Module must be configured as list or scalar")
		}
		for i := range list {
			up.AddNode(list[i])
		}
	}
}
