package main

import (
	"html/template"
	"os"
)

type test struct {
	fileSpace string
}

type error interface {
	Error() string
}

func ifErr(error interface{}) {
	if error != nil {
		//add what ever here for what do do in error
		//not graceful Demonstrates what I was saying in class
		panic(error)
	}
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
	ifErr(err)
	// if err != nil {
	// 	panic(err)
	// }

}
