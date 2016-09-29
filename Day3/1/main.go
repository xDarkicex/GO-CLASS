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
	fmt.Println("Listen on: " + "0.0.0.0" + ":8080")

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
		io.WriteString(c, fmt.Sprint("FROM SERVER: The Time is ", time.Now(), "\n"))
		bs, err := ioutil.ReadAll(c)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(bs))
		c.Close()
	}
}
