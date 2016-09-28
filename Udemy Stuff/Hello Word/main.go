package main

import "fmt"

const (
	name       = "Gentry"
	programmed = true
	lang       = "Go"
	love       = "❤"
)

type person struct {
	name string
}
type favLanguage struct {
	programmer       bool
	favoriteLanguage string
	passion          string
	person
}

func main() {
	hello()
}

func hello() {
	me := favLanguage{
		person: person{
			name: name,
		},
		programmer:       programmed,
		favoriteLanguage: lang,
	}

	fmt.Print("Hello, world. \nMy name is " + me.name + ".\n")
	if me.programmer == true {
		fmt.Print("I am a programmer.\n")
	}
	if me.favoriteLanguage == "Go" {
		fmt.Print("I  " + me.passion + me.favoriteLanguage + "\n")
	}
}
