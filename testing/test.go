// GO DOCUMENTATION:
// This is  file for testing new go methods
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	console := bufio.NewReader(os.Stdin)
	fmt.Println("How Do you feel?")
	fmt.Print("Enter text: ")
	text, _ := console.ReadString('\n')
	fmt.Println("Why do you feel " + text + "?")
}
