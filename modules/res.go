package modules

import "net/http"

type Res struct {
	Req    *Req
	writer http.ResponseWriter
}
