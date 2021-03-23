package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	// this is a comment
	fmt.Println("Hello, my name is Alex")

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
