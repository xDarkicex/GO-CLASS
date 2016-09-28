package main

import s "github.com/xDarkicex/GO-CLASS/lazy"

func main() {
	s.Say("Hello, " + s.Teacher.Name)
	s.Say("I have Objectified you See (s.Teacher.Name)")
	s.Say("( ͡° ͜ʖ ͡°)")
	s.Say(s.Friends[1], "Of course its my best Friend")

	mapExp := map[int]string{1: s.Friends[0], 2: s.Friends[1], 3: s.Friends[2]}
	s.Say(mapExp)

	for key := range mapExp {
		s.Say(key, "This is a Key")
	}
	for key, value := range mapExp {
		s.Say(key, value)
	}
}
