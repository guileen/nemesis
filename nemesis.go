package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"git.coding.net/leeen/thel_reverse_proxy/modules"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("Err to get working dir", err)
	}
	log.Println("dir", dir)

	freader, err := os.Open(dir + "/conf/config.yaml")
	if err != nil {
		log.Fatal("Error to load", err)
	}
	node, err := modules.Parse(freader)
	if err != nil {
		log.Fatal("Error to parse", err)
	}
	// 	log.Println(node)
	rootModule := modules.NewRootModule(node)
	log.Println("RootModule", rootModule)
	rootModule.Run()
	for {
		<-time.After(10 * time.Second)
	}
}
