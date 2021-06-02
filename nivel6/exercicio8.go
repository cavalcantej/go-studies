package main

import "fmt"

func retornafunc() func(){
	return func(){
		fmt.Println("funcao dentro de outra")
	}
}

func main(){
	x := retornafunc()
	x()
}
