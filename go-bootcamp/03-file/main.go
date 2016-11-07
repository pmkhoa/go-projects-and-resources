package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// f, err := os.Open("test.txt")
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("Cannot find file")
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Cannot find file")
	}

	fmt.Println(string(bs))
}
