package main

import "fmt"

type pessoa struct{
	nome 		string 
	sobrenome 	string
	idade 		int
}

func (p pessoa) falanome(){
	fmt.Printf("Meu nome completo é %v %v\n", p.nome, p.sobrenome)
}

func main() {
	maria := pessoa{"Maria","Augusta", 25}

	maria.falanome()
}