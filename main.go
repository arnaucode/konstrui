package main

import (
	"fmt"
	"os"
)

const rawFolderPath = "./webInput"
const newFolderPath = "./webOutput"
const konstruiConfigFile = "konstruiConfig.json"

func main() {
	c.Green("getting files from /webInput")
	c.Green("getting conifg from file konstruiConfig.json")

	//READ CONFIG: konstruiConfig.json
	readKonstruiConfig(rawFolderPath + "/" + konstruiConfigFile)
	c.Green("configuration:")
	fmt.Println(konstruiConfig)
	c.Green("templating")

	//create directory webOutput
	_ = os.Mkdir("webOutput", os.ModePerm)

	//DO TEMPLATING
	startTemplating(rawFolderPath, newFolderPath)
	c.Green("webpage finished, files at /webOutput")
}
