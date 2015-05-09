package modules

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootModule(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)
	// freader, err := os.Open("~/workset/thel/thel_reverse_proxy/conf/config.yaml")
	freader, err := os.Open("~/gowork/src/git.coding.net/leeen/thel_reverse_proxy/conf/config.yaml")
	assert.NoError(t, err)
	node, err := Parse(freader)
	assert.NoError(t, err)
	// 	log.Println(node)
	rootModule := NewRootModule(node)
	log.Println("RootModule", rootModule)
	go rootModule.Run()
}
