package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//TimmyConfig is the configuration from the file timmyConfig.json, on the folder /webInput
type TimmyConfig struct {
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	Github  string   `json:"github"`
	Website string   `json:"website"`
	Files   []string `json:"files"`
}

var timmyConfig TimmyConfig

func readTimmyConfig(path string) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &timmyConfig)
}
