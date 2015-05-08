package modules

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootModule(t *testing.T) {
	freader, err := os.Open("/Users/gl/workset/thel/thel_reverse_proxy/conf/config.yaml")
	assert.NoError(t, err)
	node, err := Parse(freader)
	assert.NoError(t, err)
	// 	log.Println(node)
	rootModule := NewRootModule(node)
	log.Println("RootModule", rootModule)
}
