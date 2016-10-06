package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

// Port for listen and serve
type server int

func (m server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/javascript;")
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Method      string
		Submissions map[string][]string
		URL         *url.URL
	}{
		req.Method,
		req.Form,
		req.URL,
	}

	tpl.ExecuteTemplate(res, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var s server
	http.ListenAndServe(":3000", s)
}
