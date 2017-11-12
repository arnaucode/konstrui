package main

import (
	"fmt"
	"os"
)

const rawFolderPath = "./webInput"
const konstruiConfigFile = "konstruiConfig.json"

func main() {
	c.Green("getting files from /webInput")
	c.Green("getting conifg from file konstruiConfig.json")

	//READ CONFIG: konstruiConfig.json
	konstruiConfigs := readKonstruiConfig(rawFolderPath + "/" + konstruiConfigFile)
	c.Green("configuration:")
	fmt.Println(konstruiConfigs)
	c.Green("templating")

	//create directory webOutput
	_ = os.Mkdir("webOutput", os.ModePerm)

	//DO TEMPLATING
	for _, konstruiConfig := range konstruiConfigs {
		_ = os.Mkdir(konstruiConfig.OutputDir, os.ModePerm)
		startTemplating(rawFolderPath, konstruiConfig.OutputDir, konstruiConfig)
	}
	c.Green("webpage finished, files at /webOutput")
}
