package main

type person struct {
    name string
}

func (p person) constructor(params ...string) person {
    name := "Unnamed Person" // Default name
    if len(params) > 0 {
        name = params[0]
    }
    return person{name: name}
}
func main() {
    // THIS WORKS
    person.constructor(person{},"Bob")
}