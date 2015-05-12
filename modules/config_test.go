package modules

import (
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
	assert.NotNil(t, node)
	assert.NoError(t, err)
	freader, err := os.Open(GetConfigPath())
	assert.NoError(t, err)
	assert.NotNil(t, reader)
	node, err = Parse(freader)
	assert.NotNil(t, node)
	assert.NoError(t, err)
}
