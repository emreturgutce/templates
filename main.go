package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", map[string]string{
		"India":   "Gandhi",
		"America": "MLK",
		"Prophet": "Muhammad",
	})

	if err != nil {
		log.Fatalln(err)
	}
}
