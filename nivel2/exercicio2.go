package main

import "fmt"

var x int = 87
var y int = 99

func main() {
	
	z := (x == y)
	fmt.Println(z)

	z = (x >= y)
	fmt.Println(z)

	z = (x <= y)
	fmt.Println(z)


	z = (x > y)
	fmt.Println(z)

	z = (x < y)
	fmt.Println(z)


}