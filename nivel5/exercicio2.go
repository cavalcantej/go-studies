package main

import "fmt"

type pessoa struct {
	nome 		string
	sobrenome 	string
	sabores		[]string
}

func main() {

	meumapa := make(map[string]pessoa)

	meumapa["Pimentão"] = pessoa{
		nome: "Renata",
		sobrenome: "Pimentão",
		sabores: []string{"pistache", "baunilha", "morango"}}

	meumapa["Camburão"] = pessoa{
		nome: "Carla",
		sobrenome: "Camburão",
		sabores: []string{"morango", "barraco", "briga"}}

	
	for _ , value := range meumapa {
		fmt.Printf("Meu nome é, %v  meu sobrenome é %v meus sorvetes favoritos são: \n", value.nome, value.sobrenome)
		for _ , v := range value.sabores {
			fmt.Println("--", v)
		}
		fmt.Printf("\n")
	}
	


}
