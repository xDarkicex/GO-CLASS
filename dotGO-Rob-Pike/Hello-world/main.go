package main

import (
	"fmt"
	"net/http"
)

//The error built-in interface type is the conventional interface
//for representing an error condition, with the nil value representing no error.
type error interface {
	ifErr() string
}

func ifErr(error interface{}) {
	if error != nil {
		//add what ever here for what do do in error
		//not graceful but will Demonstrates what I was saying in class
		fmt.Println(error)
	}
}
func main() {
	go willError()
	Handler()
}
func こんにちは(w http.ResponseWriter, r *http.Request) {
	//こんにちは, ジェントリー = Hello Gentry!
	fmt.Fprintf(w, "こんにちは, ジェントリー! \n")
}

//Handler servers hello world
func Handler() {
	http.HandleFunc("/", こんにちは)
	//err will hold error
	err := http.ListenAndServe("localhost:8080", nil /* no handler needed*/)
	//err is manually set to not nil all errors under hood are strings
	//-ROB PIKE
	ifErr(err)
	// if err != nil {
	// 	panic(err)
	// }
	//both take err that is being returned
}

func willError() {
	err := "This is test error"
	ifErr(err)
}
