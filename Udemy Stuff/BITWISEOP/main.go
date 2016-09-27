package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	_ = iota // 0
	//Kb kilobyte 1KB
	Kb = 1 << (iota * 10)
	//Mb Megobyte 1 MB
	Mb = 1 << (iota * 10)
	//Gb Gigabyte 1 GB
	Gb = 1 << (iota * 10)
	//Tb Terabye 1 TB
	Tb = 1 << (iota * 10)
)

func main() {
	//code here
	// if file isnt to large move
}

// MoveFile Moves file location persudo code
func MoveFile(x string, y string) {
	err := os.Rename( /*Should pass file locals here here */ "x", "y")
	if err != nil {
		err := errors.New("This is error")
		fmt.Println(err)
	}
}
