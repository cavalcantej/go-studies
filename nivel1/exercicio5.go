package main

import "fmt"

type ninja int

var x ninja
var y int

func main() {

	fmt.Printf("%v\n%T\n", x, x)
	x = 42

	fmt.Printf("%v\n", x)

	y = int(x)

	fmt.Printf("%v\n%T", y, y)

}