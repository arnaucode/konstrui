package main

import (
	"io/ioutil"
	"os"
	"strings"
)

const rawFolderPath = "./webInput"
const newFolderPath = "./webOutput"

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
func main() {
	c.Green("getting files from /webInput")
	c.Green("templating")
	parseDir(rawFolderPath, newFolderPath)
	c.Green("webpage finished, wiles at /webOutput")
}
