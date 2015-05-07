package core

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type CoreOpts struct {
	Addr         string
	UpstreamAddr string
}

type Core struct {
	Opts   *CoreOpts
	rproxy *httputil.ReverseProxy
}

func NewCore(opts *CoreOpts) *Core {
	core := &Core{Opts: opts}
	if opts.UpstreamAddr != "" {
		upurl, err := url.Parse(opts.UpstreamAddr)
		if err != nil {
			panic(err.Error())
		}
		core.rproxy = httputil.NewSingleHostReverseProxy(upurl)
	}
	return core
}

func (core *Core) Run() error {
	return http.ListenAndServe(core.Opts.Addr, core)
}

// ServeHTTP make Core implements http.Handler
func (core *Core) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if core.rproxy != nil {
		core.rproxy.ServeHTTP(rw, req)
		return
	}
	rw.Write([]byte("helloworld"))
}
