package main

import "fmt"

func main() {
	//nested()
	var v bool
	Nests(v)
}

//Nests demonstrates nested funct
func Nests(v bool) bool {

	for true {
		for i := 0; i < 30; i++ {
			if i == 20 {
				fmt.Println("I am Done")
				break
			}
			fmt.Println(i)

		}
		return false
	}

	return false
}
