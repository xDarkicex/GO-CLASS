package main

import (
	"fmt"
)

/*Interesting returning another function inside function
  Lets Go Deeper With Go
*/
//This is the Wapper Func that returns a func that returns a func that returns int
func wrapper() func() func() int {
	//non-exported local var
	index := 0
	//defining function return func int
	return func() func() int {
		// return func int
		return func() int {
			index++
			return index
		}
	}
}

func main() {
	// name wrapper
	Adder := wrapper()
	// loop Adder pointer memory to same func
	for i := 0; i < 20; i++ {
		fmt.Printf("%p\n", Adder())

	}
	//Dig Deeper Added Adder loop through Added int
	for i := 0; i < 20; i++ {
		Added := Adder()
		fmt.Printf("%d\n", Added())
	}
}
