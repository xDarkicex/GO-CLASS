//Package s 怠い
//Think of package s as me being stubburn and also I write in way to many languages
//私はプログラマーです
package s

import "fmt"

//Say is Say from perl because its less typing Please note I will find a better way to handle multi args
func Say(x ...interface{}) {
	for _, line := range x {
		fmt.Println(line)
	}
}

//Person export
type Person struct {
	Name string
	Talk func(...interface{})
}

//Name of teacher This also a struct literal
var (
	Talk    = func(x ...interface{}) { fmt.Println(x) }
	Teacher = Person{"Todd", Talk}
	Friend  = Person{"Caaz", Talk}
	Friend2 = Person{"Nina", Talk}
	Friends = []string{Teacher.Name, Friend.Name, Friend2.Name}
)
