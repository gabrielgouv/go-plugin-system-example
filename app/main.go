package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"plugin"
)

const pluginsRootDir = "./plugins"

func main() {

	files, err := ioutil.ReadDir(pluginsRootDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println("Loading plugin: " + f.Name())

		p, err := plugin.Open(pluginsRootDir + "/" + f.Name() + "/main.so")
		if err != nil {
			log.Println("Could not load " + f.Name())
		}

		pluginName, err := p.Lookup("Name")
		if err != nil {
			panic(err)
		}

		pluginVersion, err := p.Lookup("Version")
		if err != nil {
			panic(err)
		}

		pluginLoadFunc, err := p.Lookup("Load")
		if err != nil {
			panic(err)
		}

		log.Println("Success: " + *pluginName.(*string) + " - " + *pluginVersion.(*string))

		pluginLoadFunc.(func())()
	}

}
