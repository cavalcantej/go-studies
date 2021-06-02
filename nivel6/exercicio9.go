package main

import "fmt"

func argfunc(){
	fmt.Println("função argumento")
}

func funcaofuncao(x func()){
	fmt.Println("funcão que recebe função")

	x()

}

func main(){
	funcaofuncao(argfunc)
}