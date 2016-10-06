package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "test2.txt"
	newPath := "NewFile.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
