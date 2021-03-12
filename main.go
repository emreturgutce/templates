package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	})

	if err != nil {
		log.Fatalln(err)
	}
}
