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
	Files       []string `json:"files"`
	RepeatPages []RepeatPages
}

var konstruiConfig KonstruiConfig

func readKonstruiConfig(path string) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &konstruiConfig)
}
