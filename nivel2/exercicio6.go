package main 

import (
	"fmt"
)

const (
	_ = 2021 + iota
	x
	y
	z
	w
)

func main() {

	fmt.Printf("%d\n", x)
	fmt.Printf("%d\n", y)
	fmt.Printf("%d\n", z)
	fmt.Printf("%d\n", w)

}