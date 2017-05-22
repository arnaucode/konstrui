package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const rawFolderPath = "./webInput"
const newFolderPath = "./webOutput"
const konstruiConfigFile = "konstruiConfig.json"

func parseDir(folderPath string, newDir string) {
	files, _ := ioutil.ReadDir(folderPath)
	for _, f := range files {
		fileNameSplitted := strings.Split(f.Name(), ".")
		extension := fileNameSplitted[len(fileNameSplitted)-1]
		if extension == "html" {
			fileContent := putTemplates(folderPath, f.Name())
			writeFile(newDir+"/"+f.Name(), fileContent)
		} else if extension == "css" {
			fileContent := readFile(folderPath, f.Name())
			writeFile(newDir+"/"+f.Name(), fileContent)
		}
		if len(fileNameSplitted) == 1 {
			newDir := newDir + "/" + f.Name()
			oldDir := rawFolderPath + "/" + f.Name()
			if _, err := os.Stat(newDir); os.IsNotExist(err) {
				_ = os.Mkdir(newDir, 0700)
			}
			parseDir(oldDir, newDir)
		}
	}
}
func startTemplating(folderPath string, newDir string) {
	for i := 0; i < len(konstruiConfig.Files); i++ {
		fName := konstruiConfig.Files[i]
		fileNameSplitted := strings.Split(fName, ".")
		extension := fileNameSplitted[len(fileNameSplitted)-1]
		if extension == "html" {
			fileContent := putTemplates(folderPath, fName)
			writeFile(newDir+"/"+fName, fileContent)
		} else if extension == "css" {
			fileContent := readFile(folderPath, fName)
			writeFile(newDir+"/"+fName, fileContent)
		}
	}
}
func main() {
	c.Green("getting files from /webInput")
	c.Green("getting conifg from file konstruiConfig.json")
	readKonstruiConfig(rawFolderPath + "/" + konstruiConfigFile)
	c.Green("configuration:")
	fmt.Println(konstruiConfig.Files)
	c.Green("templating")
	//parseDir(rawFolderPath, newFolderPath)
	startTemplating(rawFolderPath, newFolderPath)
	c.Green("webpage finished, wiles at /webOutput")
}
