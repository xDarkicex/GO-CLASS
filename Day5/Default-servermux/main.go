package main

import (
	"io"
	"net/http"
)

type index int

func (i index) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am an index")
}

type details int

func (d details) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Me details about product")
}

type about int

func (a about) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "About ME")
}

type contact int

func (c contact) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Contact information")
}

func main() {
	var details details
	var about about
	var contact contact
	var index index

	http.Handle("/", index)
	http.Handle("/details", details)
	http.Handle("/about", about)
	http.Handle("/contact", contact)

	http.ListenAndServe(":8080", nil)
}

// change code to use DefaultServeMux
// add a route for hamsters
