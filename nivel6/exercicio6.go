package main

import "fmt"

func main(){
	o := 1000

	func(o int){
		o = 1
		fmt.Println("Função anonima")
		fmt.Println((o * 10 * 197 -2 + 544 )/4)
	}(o)
}
