//https://www.goinggo.net/2013/09/recursion-and-tail-calls-in-go_26.html
package main

import "fmt"

func main() {
	//code here
	//asking for trouble
	//ORDER MATTERS
	ch := make(chan bool)
	isRecursive(10, ch)
	//ch := make(chan bool) Dont work
	<-ch
}

//isRecursive This demostrate Recursive functions
func isRecursive(value int, ch chan bool) {
	if value < 0 {
		ch <- true
		return
	}

	fmt.Printf("%s%d\n", "Recursion value: ", value)
	//go goroutines
	go isRecursive(value-1, ch)

}

func isntRecursive() {
	//code
	fmt.Println("I am not recursive")
}
