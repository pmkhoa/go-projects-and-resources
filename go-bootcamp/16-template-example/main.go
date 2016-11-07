package main

import (
	"fmt"
	"log"
	"os"
	"html/template"
)

type Page struct {
	Title      string
	Body       string
	ScriptText template.HTML
}

func main() {
	// cwd, _ := os.Getwd()
	// templatePath := filepath.Join(cwd, "./index.gtpl")
	tmpl, err := template.ParseFiles("index.gtpl")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(os.Stdout, Page{
		Title:      "My Title",
		Body:       "<p>Hello<script>alert('Hi')</script></p>", // Template will escape this script
		ScriptText: "<p>Hello<script>alert('Hi')</script></p>",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(tmpl)
}
