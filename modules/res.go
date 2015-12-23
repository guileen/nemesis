package modules

import "net/http"

// implement http.ResponseWriter
type Res struct {
	Req    *Req
	writer http.ResponseWriter
	// is processed
	processed  bool
	statusCode int
}

// implement http.ResponseWriter
func (res *Res) Header() http.Header {
	return res.writer.Header()
}

// implement http.ResponseWriter
func (res *Res) Write(bytes []byte) (int, error) {
	if res.statusCode == 0 {
		res.WriteHeader(200)
	}
	return res.writer.Write(bytes)
}

// implement http.ResponseWriter
func (res *Res) WriteHeader(status int) {
	res.processed = true
	res.statusCode = status
	res.writer.WriteHeader(status)
}

func (res *Res) IsProcessed() bool {
	return res.processed
}
