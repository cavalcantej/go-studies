package main

import "fmt"

type pessoa struct {
	nome 		string
	sobrenome 	string
	sabores		[]string
}

func main() {
	pessoa1 := pessoa{
		nome: "Renata",
		sobrenome: "Pimentão",
		sabores: []string{"teste", "teste"}}


	fmt.Println(pessoa1)

}
