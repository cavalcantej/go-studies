package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)


	y := new(int)
	one(y)
	fmt.Println(*y)
}