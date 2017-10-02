package main

import (
	"strings"
)

func duplicateText(original string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result = result + original
	}
	return result
}
func replaceEntry(templateContent string, entry dataEntry) string {
	//replace {{}} with data
	var keys []string
	for key, _ := range entry {
		keys = append(keys, key)
	}
	//now, replace the keys with the values
	for j := 0; j < len(keys); j++ {
		templateContent = strings.Replace(templateContent, "{{"+keys[j]+"}}", entry[keys[j]], -1)
		//templateContent = strings.Replace(templateContent, "[[i]]", strconv.Itoa(i), -1)
	}
	return templateContent
}
func konstruiRepeatPartTwo(templateContent string, entries []dataEntry) string {
	var newContent string
	newContent = templateContent

	//replace <konstrui-repeat>
	if strings.Contains(newContent, "<konstrui-repeat") && strings.Contains(newContent, "</konstrui-repeat>") {
		//get content inside tags
		//get tags, and split by tags, get the content between tags
		extracted := extractText(newContent, "<konstrui-repeat", "</konstrui-repeat>")
		//for each project, putDataInTemplate data:entries, template: content inside tags

		var replaced string
		for _, entry := range entries {
			replaced = replaced + replaceEntry(extracted, entry)
		}
		fragmentLines := getLines(replaced)
		fragmentLines = deleteArrayElementsWithString(fragmentLines, "konstrui-repeat")
		//afegir fragment al newContent, substituint el fragment original
		lines := getLines(templateContent)
		p := locateStringInArray(lines, "konstrui-repeat")
		lines = deleteLinesBetween(lines, p[0], p[1])
		lines = addElementsToArrayPosition(lines, fragmentLines, p[0])
		templateContent = concatStringsWithJumps(lines)
	}
	return templateContent
}
func konstruiRepeat(templateContent string) string {
	if strings.Contains(templateContent, "<konstrui-repeat") {
		dataPath, _ := getTagParameters(templateContent, "konstrui-repeat", "repeatJSON", "nil")
		dataPath = strings.Replace(dataPath, "\n", "", -1)
		entries := getDataFromJson(rawFolderPath + "/" + dataPath)
		templateContent = konstruiRepeatPartTwo(templateContent, entries)
	}
	return templateContent
}
func konstruiSimpleVars(template string, entries []dataEntry) string {
	//now, replace simple templating variables {{vars}}
	for _, entry := range entries {
		template = replaceEntry(template, entry)
	}
	return template
}
func konstruiTemplate(templateContent string) string {
	var result string
	lines := getLines(templateContent)
	for _, line := range lines {
		if strings.Contains(line, "<konstrui-template") && strings.Contains(line, "</konstrui-template>") {
			templatePath, data := getTagParameters(line, "konstrui-template", "html", "data")
			result = result + useKonstruiTemplate(templatePath, data) + "\n"
		} else {
			result = result + line + "\n"
		}
	}
	return result
}
func useKonstruiTemplate(templatePath string, dataPath string) string {
	filepath := rawFolderPath + "/" + templatePath
	templateContent := readFile(filepath)
	entries := getDataFromJson(rawFolderPath + "/" + dataPath)
	generated := konstruiRepeatPartTwo(templateContent, entries)
	generated = konstruiSimpleVars(generated, entries)
	return generated
}
func getTagParameters(line string, tagname string, param1 string, param2 string) (string, string) {
	var param1content string
	var param2content string
	line = strings.Replace(line, "<"+tagname+" ", "", -1)
	line = strings.Replace(line, "></"+tagname+">", "", -1)
	attributes := strings.Split(line, " ")
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

func startTemplating(folderPath string, newDir string) {
	//FILES
	//do templating for each file in konstruiConfig.Files
	//konstrui-template
	for i := 0; i < len(konstruiConfig.Files); i++ {
		fName := konstruiConfig.Files[i]
		fileContent := readFile(folderPath + "/" + fName)
		fileContent = konstruiTemplate(fileContent)
		generatedPage := konstruiRepeat(fileContent)
		writeFile(newDir+"/"+fName, generatedPage)
	}
	//REPEATPAGES
	//do templating for the file pages in konstruiConfig.RepeatPages
	c.Cyan("starting to generate Pages to repeat")
	for i := 0; i < len(konstruiConfig.RepeatPages); i++ {
		pageTemplate, data := getHtmlAndDataFromRepeatPages(konstruiConfig.RepeatPages[i])
		for j := 0; j < len(data); j++ {
			var dataArray []dataEntry
			dataArray = append(dataArray, data[j])
			generatedPage := konstruiRepeatPartTwo(pageTemplate, dataArray)
			generatedPage = konstruiSimpleVars(generatedPage, dataArray)
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
