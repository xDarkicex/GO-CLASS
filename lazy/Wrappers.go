//Package s 怠い
//Think of package s as me being stubburn and also I write in way to many languages
//私はプログラマーです
package s

import "fmt"

//Say is Say from perl because its less typing Please note I will find a better way to handle multi args
func Say(x, y interface{}) {
	fmt.Println(x, y)
}

//Person export
type Person struct {
	Name string
}

//Name of teacher This also a struct literal
var (
	Teacher = Person{"Todd"}
	Friend  = Person{"Caaz"}
	Friend2 = Person{"Nina"}
	Friends = []string{Teacher.Name, Friend.Name, Friend2.Name}
)
