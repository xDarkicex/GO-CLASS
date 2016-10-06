package main

import "github.com/xDarkicex/Say"

func main() {
	fizz()
}

func fizz() {
	for i := 0; i <= 1000; i++ {
		if i%15 == 0 {
			Say.Say(i, " -- FizzBuzz \n")
		} else if i%3 == 0 {
			Say.Say(i, " -- Fizz \n")
		} else if i%5 == 0 {
			Say.Say(i, " -- Buzz \n")
		} else {
			Say.Say(i)
		}
	}
}
