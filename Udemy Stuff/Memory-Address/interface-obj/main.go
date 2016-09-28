package main

import "fmt"

type address struct {
	a int
}

type this interface {
	memory() *int
}

func (ad address) memory() *int {

	/*reflect.ValueOf(&ad).Pointer() research laws of reflection */
	var b = &ad.a

	return b
}

func main() {
	// var b *int

	//code init in here
	thisAddress := address{
		a: 42,
	}
	thatAddress := address{
		a: 43,
	}
	// not sure why this doesnt return memory address as well?
	var i this
	i = thisAddress
	a := i.memory()

	fmt.Println("I am retruned", a)
	fmt.Println("I am retruned", *a)
	i = thatAddress
	c := i.memory()
	fmt.Println("I am retruned", c)
	fmt.Println("I am retruned", *c)
}
