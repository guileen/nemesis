package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/codegangsta/cli"
	"github.com/guileen/nemesis/modules"
)

func main() {

	app := cli.NewApp()
	app.Name = "nemesis"
	app.Usage = "the server"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "",
		},
	}

	log.SetFlags(log.Ltime | log.Lshortfile)

	app.Action = func(c *cli.Context) {

		configFile := c.String("config")
		log.Println("configFile", configFile)
		if configFile == "" {
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatal("Error to get current working directory", err)
			}
			// run as static server in current directory
			rootConfig := modules.Map{
				"*": modules.Map{
					"default": modules.Scalar("true"),
					"port":    modules.Scalar("8080"),
					"routes": modules.List{
						modules.Map{
							"*": modules.Map{
								"static": modules.Scalar(cwd),
							},
						},
					},
				},
			}
			root := modules.NewRootModule(rootConfig)
			root.Run()
			for {
				<-time.After(10 * time.Second)
			}
			return
		}
		configFile, err := filepath.Abs(configFile)
		if err != nil {
			log.Fatal("Error to get config file", err)
		}

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal("Err to get working dir", err)
		}
		log.Println("dir", dir)

		freader, err := os.Open(configFile)
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

	app.Run(os.Args)
}
