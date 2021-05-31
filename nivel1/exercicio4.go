package main

import "fmt"

type ninja int

var x ninja

func main() {
	x = 42

	fmt.Printf("%v\n%T", x, x)
}