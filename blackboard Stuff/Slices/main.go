package main

import (
	"bufio"
	"fmt"
	"os"
)

var s string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter New String Must be 17 char long: ")
	for scanner.Scan() {

		input := scanner.Text()

		sliceExample(string(input))
	}
}

func sliceExample(s string) {

	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))
	fmt.Println(string([]byte(s)[:5]))
	fmt.Println(string([]byte(s)[2:5]))
	fmt.Println(string([]byte(s)[5:]))

	for _, v := range []byte(s) {
		fmt.Println(string(v))
	}
	fmt.Println("Thank you for trying.")
	os.Exit(1)
}
