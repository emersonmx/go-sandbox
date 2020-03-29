package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	text := []byte("hello\nfiles\n")
	err := ioutil.WriteFile("hello", text, 0644)
	if err != nil {
		panic(err)
	}

	dat, err := ioutil.ReadFile("hello")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dat))
}
