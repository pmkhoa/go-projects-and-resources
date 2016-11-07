package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `
        {
          "key": "testing value",
          "key2": 1234
        }`
	// json data has its value can be either string or int
	// we use interface{} to handle this.
	var obj map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}

	fmt.Println(obj)

	jsonData = `"100"`

	// Using interface{} when unsure what in json
	var obj2 interface{}
	err = json.Unmarshal([]byte(jsonData), &obj2)
	if err != nil {
		panic(err)
	}

	// always check ok when doing type casting
	v, ok := obj2.(float64)
	if !ok {
		v = 0
	}
	fmt.Println(100 + v)

	// Use Encoder or Decoder when we want to read or write to file.
}
