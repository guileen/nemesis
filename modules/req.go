package modules

import "net/http"

type Req struct {
	Res     *Res
	request *http.Request
}

func (req *Req) GetPath() string {
	return req.request.URL.Path
}
