package main

import "fmt"

func retornaint() int {
	return 10
}

func retornaintstr() (int, string){
	return 10, "olá"
}

func main(){
	fmt.Println(retornaintstr())
	fmt.Println(retornaint())

}