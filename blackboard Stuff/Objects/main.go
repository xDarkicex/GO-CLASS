package objects

import "fmt"

type gator int

// func (g gator) constructor() gator {
// 	return gator(5)
// }
func (g gator) constructor(a ...int) gator {
	value := 5
	if len(a) > 0 {
		value = a[0]
	}
	return gator(value)
}
func (g gator) greeting() {
	fmt.Printf("Hello, I am a gator that's %d feet long.\n", g)
}

type flamingo bool

func (f flamingo) constructor(a ...bool) flamingo {
	living := true
	if len(a) > 0 {
		living = a[0]
	}
	return flamingo(living)
}
func (f flamingo) greeting() {
	if f {
		fmt.Println("Hello, I'm a living flamingo")
	} else {
		fmt.Println("Hello, I'm a dead flamingo")
	}
}

type animal interface {
	greeting()
}

func main() {
    gator := gator{}
    flamingo := flamingo{}
	animals := []animal{
		gator.constructor(),
		flamingo.constructor(true),
		flamingo.constructor(false),
		//This should be alive
		flamingo.constructor(false),
		//This should be a 5 foot gator
		gator.constructor(1),
	}
	for _, currentAnimal := range animals {

		currentAnimal.greeting()
	}
}
