package modules

import (
	"io"

	"github.com/kylelemons/go-gypsy/yaml"
)

type Node interface{}
type Scalar string
type Map map[string]Node
type List []Node

func YamlNodeToNode(node yaml.Node) Node {
	scalar, ok := node.(yaml.Scalar)
	if ok {
		return Scalar(scalar.String())
	}
	list, ok := node.(yaml.List)
	if ok {
		result := make(List, list.Len())
		for i, n := range list {
			result[i] = YamlNodeToNode(n)
		}
		return result
	}
	mp, ok := node.(yaml.Map)
	if ok {
		result := make(Map)
		for k, v := range mp {
			result[k] = YamlNodeToNode(v)
		}
		return result
	}
	return nil
}

func Parse(r io.Reader) (Node, error) {
	n, err := yaml.Parse(r)
	if err != nil {
		return nil, err
	}
	return YamlNodeToNode(n), err
}

func Child(root Node, spec string) (Node, error) {
	return nil, nil
}

func (s Scalar) String() string {
	return string(s)
}

func (s Scalar) GetBool() bool {
	if s == "true" || s == "yes" || s == "1" {
		return true
	} else if s == "false" || s == "no" || s == "0" {
		return false
	}
	panic("Invalid bool value:" + s.String())
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
