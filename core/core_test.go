package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeBackendServer(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "world")
	})
	go http.ListenAndServe("localhost:5000", server)
	<-time.After(10 * time.Millisecond)
	response, err := http.Get("http://127.0.0.1:5000/hello")
	log.Println("response", response)
	assert.NoError(t, err)
}

func TestCore(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Lmicroseconds)
	makeBackendServer(t)

	core := NewCore(&CoreOpts{
		Addr:         "127.0.0.1:4000",
		UpstreamAddr: "http://127.0.0.1:5000",
	})
	go core.Run()
	response, err := http.Get("http://127.0.0.1:4000/hello")
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(response.Body)
	log.Println("response", string(body))
	assert.NoError(t, err)
}
