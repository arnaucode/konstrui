package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//RepeatPages is from the json config, is an array inside KonstruiConfig
type RepeatPages struct {
	HtmlPage string `json:"htmlPage"`
	Data     string `json:"data"`
}

//KonstruiConfig is the configuration from the file konstruiConfig.json, on the folder /webInput
type KonstruiConfig struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Github      string   `json:"github"`
	Website     string   `json:"website"`
	OutputDir   string   `json:"outputDir"`
	LangPrefix  string   `json:"langPrefix"`
	Files       []string `json:"files"`
	RepeatPages []RepeatPages
	CopyRaw     []string `json:"copyRaw"`
}

func readKonstruiConfig(path string) []KonstruiConfig {
	var konstruiConfigs []KonstruiConfig

	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &konstruiConfigs)

	return konstruiConfigs
}
