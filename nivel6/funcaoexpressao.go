package main

import "fmt"

func main(){
	x := 123


	// funções como uma expressão são atribuidas a uma variável
	y := func (x int) {
		fmt.Println(x, "vezes 87 é: ")
		fmt.Println(x * 87)
	}

	y(x)
}