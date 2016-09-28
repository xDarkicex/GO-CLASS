package main

import "fmt"

//Test from java class in go, because java sucks and go rocks
//can you tell Ive been converted?

func main() {
	outOfBounds()

}

func outOfBounds() {
	bounds := []string{"Testing", "This is one", "5", "testing", "hello"}
	fmt.Println(bounds[9])
}
