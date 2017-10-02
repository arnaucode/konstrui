package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
)

//dataEntry is the map used to create the array of maps, where the templatejson data is stored
type dataEntry map[string]string

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

func readFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	check(err)
	return string(dat)
}

func getDataFromJson(path string) []dataEntry {
	var entries []dataEntry
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	var rawEntries []*json.RawMessage
	json.Unmarshal([]byte(content), &rawEntries)
	for i := 0; i < len(rawEntries); i++ {
		rawEntryMarshaled, err := json.Marshal(rawEntries[i])
		check(err)
		var newDataEntry map[string]string
		json.Unmarshal(rawEntryMarshaled, &newDataEntry)
		entries = append(entries, newDataEntry)
	}

	return entries
}

func writeFile(path string, newContent string) {
	err := ioutil.WriteFile(path, []byte(newContent), 0644)
	check(err)

	color.Green(path + ":")
	color.Blue(newContent)
}

/*func generatePageFromTemplateAndData(templateContent string, entry dataEntry) string {
	var entryContent string
	entryContent = templateContent
	//first, get the map keys
	var keys []string
	for key, _ := range entry {
		keys = append(keys, key)
	}
	//now, replace the keys with the values
	for j := 0; j < len(keys); j++ {
		entryContent = strings.Replace(entryContent, "{{"+keys[j]+"}}", entry[keys[j]], -1)
	}
	return entryContent
}*/
func getHtmlAndDataFromRepeatPages(page RepeatPages) (string, []dataEntry) {
	templateContent := readFile(rawFolderPath + "/" + page.HtmlPage)
	data := getDataFromJson(rawFolderPath + "/" + page.Data)
	return templateContent, data
}
