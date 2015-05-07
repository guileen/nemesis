package config

import (
	"log"
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
}
