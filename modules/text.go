package modules

import (
	"fmt"
	"log"
)

type TextModule struct {
	BaseModule
	text []byte
}

func NewTextModule(node Node) *TextModule {
	m := &TextModule{}
	m.Init(node)
	return m
}

func (m *TextModule) Init(node Node) {
	scalar, ok := node.(Scalar)
	if !ok {
		panic("Bad text config:" + fmt.Sprintf("%b", node))
	}
	m.text = []byte(scalar.String())
}

func (m *TextModule) Process(req *Req, res *Res) bool {
	log.Println("Process text")
	res.Write(m.text)
	return true
}
