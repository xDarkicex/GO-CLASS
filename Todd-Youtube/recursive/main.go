//https://www.goinggo.net/2013/09/recursion-and-tail-calls-in-go_26.html
package main

import "fmt"

func main() {
	//code here

}

//isRecursive This demostrate Recursive functions
func isRecursive() {
	//code
	fmt.Println("I am recursive")
	isRecursive()
}

func isntRecursive() {
	//code
	fmt.Println("I am not recursive")
}
