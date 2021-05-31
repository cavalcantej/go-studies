package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	slice2 := slice[0:3]
	fmt.Println(slice2)
	
	slice3 := slice[4:]
	fmt.Println(slice3)

	slice4 := slice[1:7]
	fmt.Println(slice4)

	slice5 := slice[2:len(slice) - 1]
	fmt.Println(slice5)

}