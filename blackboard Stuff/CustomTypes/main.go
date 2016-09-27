package main

import "fmt"

type gator int

func (f gator) greeting() {
	fmt.Println("Hello, I am a gator, my value is " + string(f))
}

func main() {
	var g1 gator
	g1.greeting()
}
