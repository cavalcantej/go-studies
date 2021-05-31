package main

import "fmt"

func main(){
	x := 123


	// funções anonimas são chamadas no final da sua declaração. é uma função descartável
	func (x int) {
		fmt.Println(x, "vezes 87 é: ")
		fmt.Println(x * 87)
	}(x)
}