package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
func duplicateText(original string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result = result + original
	}
	return result
}
func replaceEntry(templateContent string, entry dataEntry) string {
	//var newContent string

	//replace {{}} for data
	//for i := 0; i < len(entries); i++ {
	//var entryContent string
	//entryContent = templateContent
	//first, get the map keys
	var keys []string
	for key, _ := range entry {
		keys = append(keys, key)
	}
	//now, replace the keys with the values
	for j := 0; j < len(keys); j++ {
		templateContent = strings.Replace(templateContent, "{{"+keys[j]+"}}", entry[keys[j]], -1)
		//templateContent = strings.Replace(templateContent, "[[i]]", strconv.Itoa(i), -1)
	}
	//newContent = newContent + entryContent
	//newContent = newContent + "\n"
	//}
	return templateContent
}
func putDataInTemplate(templateContent string, entries []dataEntry) string {
	var newContent string
	newContent = templateContent

	//replace <konstrui-repeat>
	if strings.Contains(newContent, "<konstrui-repeat") && strings.Contains(newContent, "</konstrui-repeat>") {
		//repeat, _ := getTagParameters(newContent, "konstrui-repeat", "repeat", "nil")

		//get content inside tags
		//get tags, and split by tags, get the content between tags
		extracted := extractText(newContent, "<konstrui-repeat", "</konstrui-repeat>")
		//for each project, putDataInTemplate data:entries, template: content inside tags

		//var fragment string
		var replaced string
		for _, entry := range entries {
			color.Green(extracted)
			fmt.Println(entry)
			replaced = replaced + replaceEntry(extracted, entry)
		}
		fragmentLines := getLines(replaced)
		fragmentLines = deleteArrayElementsWithString(fragmentLines, "konstrui-repeat")
		//afegir fragment al newContent, substituint el fragment original
		lines := getLines(templateContent)
		p := locateStringInArray(lines, "konstrui-repeat")
		lines = deleteLinesBetween(lines, p[0], p[1])
		lines = addElementsToArrayPosition(lines, fragmentLines, p[0])
		/*lines = deleteArrayElementsWithString(lines, "konstrui-repeat")
		fmt.Println(lines)*/
		templateContent = concatStringsWithJumps(lines)
		fmt.Println(templateContent)
		newContent = templateContent
	}
	color.Red(newContent)
	result := templateContent
	for _, entry := range entries {
		result = replaceEntry(result, entry)
	}
	color.Blue(result)

	return result
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
func useTemplateContent(templateContent string, dataPath string) string {
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
			fileContent = fileContent + useTemplate(templatePath, data) + "\n"
		} else {
			fileContent = fileContent + currentLine + "\n"
		}
		lineCount++
	}

	if strings.Contains(fileContent, "<konstrui-repeat") {
		dataPath, _ := getTagParameters(fileContent, "konstrui-repeat", "repeat", "nil")
		dataPath = strings.Replace(dataPath, "\n", "", -1)
		fileContent = useTemplateContent(fileContent, dataPath)
		color.Red(fileContent)
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
