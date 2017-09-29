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

/*func parseDir(folderPath string, newDir string) {
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
			newDir = newDir + "/" + f.Name()
			oldDir := rawFolderPath + "/" + f.Name()
			if _, err := os.Stat(newDir); os.IsNotExist(err) {
				_ = os.Mkdir(newDir, 0700)
			}
			parseDir(oldDir, newDir)
		}
	}
}*/
func startTemplating(folderPath string, newDir string) {
	//FILES
	//do templating for each file in konstruiConfig.Files in konstruiConfig.Files
	//konstrui-template
	for i := 0; i < len(konstruiConfig.Files); i++ {
		fName := konstruiConfig.Files[i]
		//fileNameSplitted := strings.Split(fName, ".")
		//extension := fileNameSplitted[len(fileNameSplitted)-1]
		fileContent := putTemplates(folderPath, fName)
		/*fmt.Println(i)
		color.Red(fileContent)*/
		writeFile(newDir+"/"+fName, fileContent)
	}
	//REPEATPAGES
	//do templating for the file pages in konstruiConfig.RepeatPages
	c.Cyan("starting to generate Pages to repeat")
	for i := 0; i < len(konstruiConfig.RepeatPages); i++ {
		pageTemplate, data := getHtmlAndDataFromRepeatPages(konstruiConfig.RepeatPages[i])
		for j := 0; j < len(data); j++ {
			//fmt.Println(j)
			var dataArray []dataEntry
			dataArray = append(dataArray, data[j])
			generatedPage := putDataInTemplate(pageTemplate, dataArray)
			//fmt.Println(data[j])
			writeFile(newDir+"/"+data[j]["pageName"]+"Page.html", generatedPage)
		}
	}
	//COPYRAW
	//copy the konstruiConfig.CopyRaw files without modificate them
	for i := 0; i < len(konstruiConfig.CopyRaw); i++ {
		fName := konstruiConfig.CopyRaw[i]
		c.Yellow(fName)
		fileNameSplitted := strings.Split(fName, ".")
		if len(fileNameSplitted) > 1 {
			//is a file
			copyFileRaw(folderPath, fName, newDir)
		} else {
			//is a directory
			c.Red(folderPath + "/" + fName)
			copyDirRaw(folderPath, fName, newDir)
		}
	}
}
func copyDirRaw(fromDir string, currentDir string, newDir string) {
	filesList, _ := ioutil.ReadDir("./" + fromDir + "/" + currentDir)
	fmt.Println(fromDir + "/" + currentDir)
	c.Green(newDir + "/" + currentDir)
	os.MkdirAll(newDir+"/"+currentDir, os.ModePerm)
	for _, f := range filesList {
		fileNameSplitted := strings.Split(f.Name(), ".")
		if len(fileNameSplitted) > 1 {
			//is a file
			copyFileRaw(fromDir+"/"+currentDir, f.Name(), newDir+"/"+currentDir)
		} else {
			//is a directory
			copyDirRaw(fromDir+"/"+currentDir, f.Name(), newDir+"/"+currentDir)
		}
	}
}
func copyFileRaw(fromDir string, fName string, newDir string) {
	c.Yellow("copying raw " + fromDir + "/" + fName)
	fileContent := readFile(fromDir + "/" + fName)
	writeFile(newDir+"/"+fName, fileContent)
}
func main() {
	c.Green("getting files from /webInput")
	c.Green("getting conifg from file konstruiConfig.json")
	//first reads the konstrui.Config.json
	readKonstruiConfig(rawFolderPath + "/" + konstruiConfigFile)
	c.Green("configuration:")
	fmt.Println(konstruiConfig)
	c.Green("templating")
	//parseDir(rawFolderPath, newFolderPath)

	//create directory webOutput
	_ = os.Mkdir("webOutput", os.ModePerm)

	startTemplating(rawFolderPath, newFolderPath)
	c.Green("webpage finished, files at /webOutput")
}
