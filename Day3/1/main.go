package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		c, err := li.Accept()
		if err != nil {
			panic(err)
		}
		bs, err := ioutil.ReadAll(c)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(bs))

		io.WriteString(c, fmt.Sprint("FROM SERVER: HELLO CLASS ", time.Now(), "\n"))

		c.Close()
	}
}
