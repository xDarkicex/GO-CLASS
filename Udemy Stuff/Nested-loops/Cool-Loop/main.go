package main

import "github.com/xDarkicex/Say"

func init() {

}
func main() {
	//Start code here
	mod()
}

// //Say stuff
// //Say is Say from perl because its less typing Please note I will find a better way to handle multi args
// func Say(x ...interface{}) {
// 	for _, line := range x {
// 		fmt.Print(line)
// 	}
// }

func mod() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		Say.Say(i, " Is a Cool Cat!\n")
		if i >= 50 {
			break
		}
	}

}
