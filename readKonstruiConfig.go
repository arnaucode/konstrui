package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//KonstruiConfig is the configuration from the file konstruiConfig.json, on the folder /webInput
type KonstruiConfig struct {
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	Github  string   `json:"github"`
	Website string   `json:"website"`
	Files   []string `json:"files"`
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
