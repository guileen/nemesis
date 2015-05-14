package modules

import (
	"fmt"
	"log"
)

type TextModule struct {
	text []byte
}

func NewTextModule(node Node) *TextModule {
	m := &TextModule{}
	m.Init(node)
	return m
}

func (m *TextModule) Init(node Node) {
	txt, ok := node.(Scalar).String()
	if !ok {
		panic("Bad text config:" + fmt.Sprintf("%b", node))
	}
	m.text = txt
}

func (m *TextModule) Process(req *Req, res *Res) bool {
	log.Println("Process text")
	res.writer.Write(m.text)
	return true
}
