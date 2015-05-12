package modules

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRootModule(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)
	// freader, err := os.Open("~/workset/thel/thel_reverse_proxy/conf/config.yaml")
	freader, err := os.Open(GetConfigPath())
	assert.NoError(t, err)
	node, err := Parse(freader)
	assert.NoError(t, err)
	// 	log.Println(node)
	rootModule := NewRootModule(node)
	log.Println("RootModule", rootModule)
	go rootModule.Run()
	<-time.After(10 * time.Millisecond)
	resp, err := http.Get("http://127.0.0.1:8000/hello")
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	log.Println("resp", string(body))
}
