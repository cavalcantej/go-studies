package main

import "fmt"


const x int = 10

const y = 10.5

func main() {

	fmt.Printf("%v\t%T\n", x, x)
	fmt.Printf("%v\t%T", y, y)
}