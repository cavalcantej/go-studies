package main 

import "fmt"

var x int

func main() {
	x = 10
	fmt.Printf("%d\t%b\t%x\n", x , x, x)

	y := x << 1
	fmt.Printf("%d\t%b\t%x\n", y, y, y)
}