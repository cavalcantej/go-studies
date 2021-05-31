package main

import "fmt"

func main() {
	
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50}

	y := []int{}
	y = append(x[0:3], x[6:]...)
	y = append(y, 51)
	
	fmt.Println(y)


}