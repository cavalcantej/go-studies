package main

import "fmt"

func retornaint() int {
	return 10
}

func imprime() {
	fmt.Println("Olá mundo")
}

func main(){
	defer imprime()
	
	fmt.Println(retornaint())



}