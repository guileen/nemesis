package modules

import (
	"encoding/json"
	"fmt"
	"log"
)

type JSONModule struct {
	BaseModule
	text []byte
}

func NewJSONModule(node Node) *JSONModule {
	m := &JSONModule{}
	m.Init(node)
	return m
}

func (m *JSONModule) Init(node Node) {
	txt, err := json.Marshal(node)
	if err != nil {
		panic("Bad json config:" + fmt.Sprintf("%b", node))
	}
	m.text = txt
}

func (m *JSONModule) Process(req *Req, res *Res) bool {
	log.Println("Process json")
	res.Write(m.text)
	return true
}
