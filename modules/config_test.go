package modules

import (
	"log"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetConfigPath() string {
	wd, _ := os.Getwd()
	confPath := path.Clean(wd + "/../conf/config.yaml")
	return confPath
}

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
	freader, err := os.Open(GetConfigPath())
	assert.NoError(t, err)
	assert.NotNil(t, reader)
	node, err = Parse(freader)
	assert.NoError(t, err)
	log.Println(node)
}
