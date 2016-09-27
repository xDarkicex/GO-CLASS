package main

import (
	"bytes"

	s "github.com/xDarkicex/GO-CLASS/lazy"
)

//declares the variable buffer, which holds 256 bytes.
//The type of buffer includes its size, [256]byte.
//An array with 512 bytes would be of the distinct type [512]byte.
//The data associated with an array is just that: an array of elements.
//Schematically, our buffer looks like this in memory,

var buffer [256]byte
var slice = buffer[100:150]
var slicer = slice[5:10]
var slashPos = bytes.IndexRune(slice, '/')

func main() {
	s.Say("Hello, " + s.Teacher.Name)
	s.Teacher.Talk("Something")
	s.Say("I have Objectified you See (s.Teacher.Name)")
	s.Say("( ͡° ͜ʖ ͡°)")
	//Slice of buffer
	sliceExp := buffer[10:20]
	for i := 0; i < len(sliceExp); i++ {
		sliceExp[i] = byte(i)
	}
	//Local slice will take precedence over global slice var
	slice := []int{3, 5, 6, 2, 1, 6, 34, 54, 21, 60}
	s.Say(slice, "")
	//Lets do the challenge shall we
	for index := range slice {
		//This returns key and value but, we could do  it another way.
		s.Say(slice[index])
		s.Say(index)
	}
	//Key value pair wow how cool is that
	for index, value := range slice {
		s.Say(index, value)
	}
	//SCOPING
	//Zero value of global non exported var i []byte
	s.Say(i)
	//[]byte value of adder before Todd notice bother are the Zero value of var i
	s.Say(adder(i))
	//Add Todd
	adder([]byte(s.Teacher.Name))
	//[]byte value after Todd
	s.Say(adder(i))
	s.Say("string value of adder casting")
	s.Say(string(adder(i)))
	s.Say("feed adder lower case Todd")
	adder([]byte(s.Teacher2.Name))
	s.Say("[]byte value after todd")
	s.Say(adder(i))
	s.Say("string value of adder casting")
	s.Say(string(adder(i)))

}
func addOne(sliceExp []byte) {
	for i := range slice {
		slice[i]++
		s.Say(slice[i])
	}

}

var i []byte

func adder(x []byte) []byte {
	i = []byte(x)
	return i
}
