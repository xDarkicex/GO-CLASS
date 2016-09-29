package main

import (
	"html/template"
	"os"
)

type test struct {
	fileSpace string
}

//Tpl is template parsing thing
var templates *template.Template

func init() {
	//code goes here
	templates = template.Must(template.ParseGlob("tmp/*"))

}

func main() {
	test := map[string]interface{}{
		"fileSpace": template.HTML("<link rel='stylesheet' type='text/css' href='/assets/stylesheets/application.css'>"),
	}

	//code goes here
	err := templates.ExecuteTemplate(os.Stdout, "index.gohtml", test)
	if err != nil {
		panic(err)
	}

}
