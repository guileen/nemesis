package config

import (
	"io"

	"github.com/kylelemons/go-gypsy/yaml"
)

type Node interface{}
type Scalar string
type Map map[string]Node
type List []Node

func Parse(r io.Reader) (Node, error) {
	n, err := yaml.Parse(r)
	return n, err
}

func Child(root Node, spec string) (Node, error) {
	return nil, nil
}

func (s Scalar) String() string {
	return string(s)
}

func (m Map) Key(key string) Node {
	return m[key]
}

func (l List) Item(idx int) Node {
	return l[idx]
}

func (l List) Len() int {
	return len(l)
}
