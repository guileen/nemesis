package modules

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	reader := strings.NewReader(`
  foo: bar
  map:
    a:b
    c:d
  list :
    - a
    - b
`)
	node, err := Parse(reader)
	log.Println(node)
	assert.NoError(t, err)
	// freader, err := os.Open("/Users/gl/workset/thel/thel_reverse_proxy/conf/config.yaml")
	freader, err := os.Open("/Users/gl/gowork/src/git.coding.net/leeen/thel_reverse_proxy/conf/config.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, reader)
	node, err = Parse(freader)
	assert.NoError(t, err)
	log.Println(node)
}
