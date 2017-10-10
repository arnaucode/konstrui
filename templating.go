package main

import (
	"strings"

	"github.com/Jeffail/gabs"
)

func duplicateText(original string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result = result + original
	}
	return result
}

/*func replaceEntryOLD(templateContent string, entry dataEntry, jsonData *gabs.Container, elemName string) string {
	//replace {{}} with data
	var keys []string
	for key, _ := range entry {
		keys = append(keys, key)
	}
	//now, replace the keys with the values
	for j := 0; j < len(keys); j++ {
		templateContent = strings.Replace(templateContent, "{{"+keys[j]+"}}", entry[keys[j]], -1)
		//templateContent = strings.Replace(templateContent, "[[i]]", strconv.Itoa(i), -1)
		children, _ := jsonData.S(keys[j]).Children()
		fmt.Println("-")
		fmt.Println(keys[j])
		fmt.Println(children)
		for key, child := range children {
			fmt.Print("key: " + strconv.Itoa(key) + ", value: ")
			fmt.Println(child.Data().(string))
		}
	}

	return templateContent
}*/
func replaceEntry(templateContent string, entry dataEntry, jsonData *gabs.Container, elemName string) string {
	//fmt.Println(jsonData)
	children, _ := jsonData.S().ChildrenMap()
	//_, ok := jsonData.S().Children()
	for parameter, child := range children {
		//subchildren, _ := child.S().ChildrenMap()
		_, ok := child.S().Children()
		if ok != nil {
			templateContent = strings.Replace(templateContent, "{{"+parameter+"}}", child.Data().(string), -1)
		} else {
			/*for subparameter, subchild := range subchildren {
				color.Red(subchild.Data().(string))
				fmt.Println(subchild)
				fmt.Println(subparameter)
				//templateContent = strings.Replace(templateContent, "{{"+subparameter+"}}", subchild.Data().(string), -1)
			}*/
		}

	}
	return templateContent
}
func konstruiRepeatJSONPartTwo(templateContent string, entries []dataEntry, jsonData *gabs.Container, elemName string) string {
	var newContent string
	newContent = templateContent

	//replace <konstrui-repeatJSON>
	if strings.Contains(newContent, "<konstrui-repeatJSON") && strings.Contains(newContent, "</konstrui-repeatJSON>") {
		//get content inside tags
		//get tags, and split by tags, get the content between tags
		extracted := extractText(newContent, "<konstrui-repeatJSON", "</konstrui-repeatJSON>")
		//for each project, putDataInTemplate data:entries, template: content inside tags
		var replaced string
		//for _, entry := range entries {
		children, _ := jsonData.S().Children()
		for _, child := range children {
			var entry dataEntry
			replaced = replaced + replaceEntry(extracted, entry, child, elemName)
		}
		fragmentLines := getLines(replaced)
		fragmentLines = deleteArrayElementsWithString(fragmentLines, "konstrui-repeatJSON")
		//afegir fragment al newContent, substituint el fragment original
		lines := getLines(templateContent)
		p := locateStringInArray(lines, "konstrui-repeatJSON")
		lines = deleteLinesBetween(lines, p[0], p[1])
		lines = addElementsToArrayPosition(lines, fragmentLines, p[0])
		templateContent = concatStringsWithJumps(lines)
	}

	return templateContent
}
func konstruiRepeatJSON(templateContent string) string {
	if strings.Contains(templateContent, "<konstrui-repeatJSON") {
		dataPath, _ := getTagParameters(templateContent, "konstrui-repeatJSON", "repeatJSON", "nil")
		dataPath = strings.Replace(dataPath, "\n", "", -1)
		entries, jsonData := getDataFromJson(rawFolderPath + "/" + dataPath)
		templateContent = konstruiRepeatJSONPartTwo(templateContent, entries, jsonData, "")
	}
	return templateContent
}
func konstruiRepeatElem(templateContent string, entry dataEntry, jsonData *gabs.Container) string {
	if strings.Contains(templateContent, "<konstrui-repeatElem") {
		elemName, _ := getTagParameters(templateContent, "konstrui-repeatElem", "repeatElem", "nil")
		extracted := extractText(templateContent, "<konstrui-repeatElem", "</konstrui-repeatElem>")
		fragmentLines := getLines(extracted)
		fragmentLines = deleteArrayElementsWithString(fragmentLines, "konstrui-repeatElem")
		f := concatStringsWithJumps(fragmentLines)
		children, _ := jsonData.S(elemName).Children()
		var replaced string
		for _, child := range children {
			//fmt.Println(child.Data().(string))
			replacedElem := strings.Replace(f, "{{"+elemName+"}}", child.Data().(string), -1)
			replaced = replaced + replacedElem
		}
		fragmentLines = getLines(replaced)

		lines := getLines(templateContent)
		p := locateStringInArray(lines, "konstrui-repeatElem")
		lines = deleteLinesBetween(lines, p[0], p[1])
		lines = addElementsToArrayPosition(lines, fragmentLines, p[0])
		templateContent = concatStringsWithJumps(lines)

	}
	return templateContent
}

func konstruiSimpleVars(template string, entries []dataEntry, jsonData *gabs.Container) string {
	//now, replace simple templating variables {{vars}}
	/*for _, entry := range entries {
		template = replaceEntry(template, entry, jsonData, "")
	}*/
	/*children, _ := jsonData.S().Children()
	for _, child := range children {
		var entry dataEntry
		fmt.Println("aaaaa")
		fmt.Println(child)
		template = replaceEntry(template, entry, child, "")
	}
	*/
	var entry dataEntry
	template = replaceEntry(template, entry, jsonData, "")
	return template
}
func konstruiInclude(content string) string {
	var result string
	if strings.Contains(content, "<konstrui-include") {
		lines := getLines(content)
		for _, line := range lines {
			if strings.Contains(line, "<konstrui-include") {
				dataPath, _ := getTagParameters(line, "<konstrui-include", "html", "nil")
				dataPath = strings.Replace(dataPath, "\n", "", -1)
				htmlInclude := readFile(rawFolderPath + "/" + dataPath)
				result = result + htmlInclude
			} else {
				result = result + line + "\n"
			}
		}

		/*dataPath, _ := getTagParameters(content, "<konstrui-include", "html", "nil")
		dataPath = strings.Replace(dataPath, "\n", "", -1)
		htmlInclude := readFile(rawFolderPath + "/" + dataPath)
		htmlIncludeLines := getLines(htmlInclude)
		contentLines := getLines(content)
		p := locateStringInArray(contentLines, "<konstrui-include")
		fmt.Println(p)
		contentLines = deleteLinesBetween(contentLines, p[0], p[0])
		contentLines = addElementsToArrayPosition(contentLines, htmlIncludeLines, p[0])
		content = concatStringsWithJumps(contentLines)*/
	} else {
		result = content
	}
	return result
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
	entries, jsonData := getDataFromJson(rawFolderPath + "/" + dataPath)
	generated := konstruiRepeatJSONPartTwo(templateContent, entries, jsonData, "")
	generated = konstruiSimpleVars(generated, entries, jsonData)
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
		generatedPage := konstruiRepeatJSON(fileContent)
		generatedPage = konstruiInclude(generatedPage)
		writeFile(newDir+"/"+fName, generatedPage)
	}
	//REPEATPAGES
	//do templating for the file pages in konstruiConfig.RepeatPages
	c.Cyan("starting to generate Pages to repeat")
	for i := 0; i < len(konstruiConfig.RepeatPages); i++ {
		pageTemplate, _, jsonData := getHtmlAndDataFromRepeatPages(konstruiConfig.RepeatPages[i])
		//for j := 0; j < len(data); j++ {
		children, _ := jsonData.S().Children()
		for _, child := range children {
			var dataArray []dataEntry
			var dat dataEntry
			//dataArray = append(dataArray, data[j])
			generatedPage := konstruiRepeatJSONPartTwo(pageTemplate, dataArray, child, "")
			generatedPage = konstruiRepeatElem(generatedPage, dat, child)
			generatedPage = konstruiSimpleVars(generatedPage, dataArray, child)
			generatedPage = konstruiInclude(generatedPage)
			//writeFile(newDir+"/"+data[j]["pageName"]+"Page.html", generatedPage)
			pageName, _ := child.Path("pageName").Data().(string)
			writeFile(newDir+"/"+pageName+"Page.html", generatedPage)
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
