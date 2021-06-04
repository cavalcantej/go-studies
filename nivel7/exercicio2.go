package main

import "fmt"

type pessoa struct{
	nome 		string
	sobrenome	string
	idade 		int
}

func mudeMe(p *pessoa, n string, s string, i int){
	(*p).nome = n
	(*p).sobrenome = s
	(*p).idade = i
}


func main(){
	p1 := pessoa{"josiane", "alves", 49}
	pointerPessoa := &p1

	fmt.Println("Primeiro dado de pessoa: ",p1)
	fmt.Println()
	fmt.Println("Endereço de p1.nome:", &p1.nome)
	fmt.Println("Endereço de p1.sobrenome:", &p1.sobrenome)
	fmt.Println("Endereço de p1.idade:", &p1.idade)
	fmt.Println()
	fmt.Println("Endereço de p1:", &p1)
	fmt.Println("Valor do pointer:", pointerPessoa)
	fmt.Println()
	fmt.Println("Função mudeMe() executando...")
	// pointer como parametro, e os dados a serem alterados
	// poderia ter passado apenas &p1 direto
	mudeMe(pointerPessoa, "maria", "augusta", 25)

	fmt.Println("Novo endereço de p1: ", pointerPessoa)
	fmt.Println()
	fmt.Println("Novos valores de p1... ")
	fmt.Println("p1.nome: ", p1.nome)
	fmt.Println("p1.sobrenome: ", p1.sobrenome)
	fmt.Println("p1.idade: ", p1.idade)

}