package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

//dataEntry is the map used to create the array of maps, where the templatejson data is stored
type dataEntry map[string]string

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
func replaceEntries(templateContent string, entries []dataEntry) string {
	var newContent string

	//replace {{}} for data
	for i := 0; i < len(entries); i++ {
		var entryContent string
		entryContent = templateContent
		//first, get the map keys
		var keys []string
		for key, _ := range entries[i] {
			keys = append(keys, key)
		}
		//now, replace the keys with the values
		for j := 0; j < len(keys); j++ {
			entryContent = strings.Replace(entryContent, "{{"+keys[j]+"}}", entries[i][keys[j]], -1)
			entryContent = strings.Replace(entryContent, "[[i]]", strconv.Itoa(i), -1)
		}
		newContent = newContent + entryContent
	}
	return newContent
}
func putDataInTemplate(templateContent string, entries []dataEntry) string {
	var newContent string
	newContent = templateContent

	//replace <konstrui-repeat>
	if strings.Contains(newContent, "<konstrui-repeat") && strings.Contains(newContent, "</konstrui-repeat>") {
		//repeat, _ := getTagParameters(newContent, "konstrui-repeat", "repeat", "nil")
		color.Blue("repeat data")
		//fmt.Println(repeat)

		//get content inside tags
		//get tags, and split by tags, get the content between tags
		extracted := extractText(newContent, "<konstrui-repeat", "</konstrui-repeat>")
		fmt.Println(extracted)
		//for each project, putDataInTemplate data:entries, template: content inside tags
		fragment := replaceEntries(extracted, entries)
		color.Blue(fragment)
		//afegir fragment al newContent
		//esborrar les línies dels tags
	}

	newContent = replaceEntries(templateContent, entries)

	return newContent
}
func getTagParameters(line string, tagname string, param1 string, param2 string) (string, string) {
	var param1content string
	var param2content string
	line = strings.Replace(line, "<"+tagname+" ", "", -1)
	line = strings.Replace(line, "></"+tagname+">", "", -1)
	attributes := strings.Split(line, " ")
	//fmt.Println(attributes)
	for i := 0; i < len(attributes); i++ {
		attSplitted := strings.Split(attributes[i], "=")
		if attSplitted[0] == param1 {
			param1content = strings.Replace(attSplitted[1], `"`, "", -1)
			param1content = strings.Replace(param1content, ">", "", -1)
		}
		if attSplitted[0] == param2 {
			param2content = strings.Replace(attSplitted[1], `"`, "", -1)
			param2content = strings.Replace(param2content, ">", "", -1)
		}
	}
	return param1content, param2content
}

func useTemplate(templatePath string, dataPath string) string {
	filepath := rawFolderPath + "/" + templatePath
	templateContent := readFile(filepath)
	entries := getDataFromJson(rawFolderPath + "/" + dataPath)
	generated := putDataInTemplate(templateContent, entries)
	return generated
}

func putTemplates(folderPath string, filename string) string {
	var fileContent string
	f, err := os.Open(folderPath + "/" + filename)
	check(err)
	scanner := bufio.NewScanner(f)
	lineCount := 1
	for scanner.Scan() {
		currentLine := scanner.Text()
		if strings.Contains(currentLine, "<konstrui-template") && strings.Contains(currentLine, "</konstrui-template>") {
			templatePath, data := getTagParameters(currentLine, "konstrui-template", "html", "data")
			fileContent = fileContent + useTemplate(templatePath, data)
		} else {
			fileContent = fileContent + currentLine
		}
		lineCount++
	}
	return fileContent
}

func writeFile(path string, newContent string) {
	err := ioutil.WriteFile(path, []byte(newContent), 0644)
	check(err)
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
