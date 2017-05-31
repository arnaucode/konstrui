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
			path := folderPath + "/" + f.Name()
			fileContent := readFile(path)
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
			path := folderPath + "/" + fName
			fileContent := readFile(path)
			writeFile(newDir+"/"+fName, fileContent)
		}
	}

	c.Cyan("starting to generate Pages to repeat")
	for i := 0; i < len(konstruiConfig.RepeatPages); i++ {
		pageTemplate, data := getHtmlAndDataFromRepeatPages(konstruiConfig.RepeatPages[i])
		for j := 0; j < len(data); j++ {
			fmt.Println(j)
			generatedPage := generatePageFromTemplateAndData(pageTemplate, data[j])
			fmt.Println(data[j])
			writeFile(newDir+"/"+data[j]["pageName"]+"Page.html", generatedPage)
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
