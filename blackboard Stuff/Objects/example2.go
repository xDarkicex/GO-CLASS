package main

import "fmt"

type person struct {
	name string
}

// This is working, but a really ugly way to do so
func personConstructor(params ...string) person {
	name := "Unnamed Person" // Default name
	if len(params) > 0 {
		name = params[0]
	}
	return person{name: name}
}
func (p person) talk() {
	fmt.Printf("Hello, my name is %s\n", p.name)
}

func main() {
	bob := personConstructor("Bob")
	bob.talk()
}
