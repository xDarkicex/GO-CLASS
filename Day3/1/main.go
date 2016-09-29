package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listen on: " + "0.0.0.0" + ":8080")
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		defer li.Close()
		c, err := li.Accept()
		if err != nil {
			panic(err)
		}
		s := bufio.NewScanner(c)
		for s.Scan() {
			ln := s.Text()
			fmt.Println(ln)
		}
	}
}
