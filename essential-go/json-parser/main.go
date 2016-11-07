package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadConfig(path string) (map[string]interface{}, error) {
	var m map[string]interface{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(data, &m)
	return m, err
}

func main() {
	config, err := LoadConfig("/Users/khoapham/Code/Learning/Golang/essential-go/json-parser/config.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
