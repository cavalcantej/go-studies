package main

import "fmt"

var s string

func main() {
	s = `this is a
	raw string %d \n literal
	"""""""""""`
	fmt.Println(s)
}