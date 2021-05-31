package main

import "fmt"

func main(){
	x := retornumafuncao()

	y := x(3)

	fmt.Println(y)
	fmt.Println(retornumafuncao()(5))
}

func retornumafuncao() func(int) int {
	return func(i int) int {
		return i * 5
	} 
}