package main

import "fmt"



type cliente struct {
	nome		string
	sobrenome	string
	fumante 	bool
}

type pessoa struct {
	nome	string
	idade	int
}

type profissao struct {
	pessoa
	titulo	string
	salario	int
}


func main() {
	cliente1 := cliente {
		nome : "Calinhos",
		sobrenome : "Bala",
		fumante : true,
	}


	cliente2 := cliente { "Lana", "Almeida", true }


	pessoa1 := pessoa{
		nome: "Tia",
		idade: 272,
	}

	pessoa2 := profissao{
		pessoa : pessoa{
			nome: "Maria",
			idade: 514,
		},
		titulo: "Recepcionista",
		salario: 2000,
	}

	pessoa3 := profissao{pessoa{"Zew",14},"Gamer", 10}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2.pessoa.nome)


	fmt.Println(pessoa3)


	fmt.Println(cliente1)
	fmt.Println(cliente2)


	x := struct{
		nome string
		idade int
	}{
		nome: "Mario",
		idade: 18}
	
	fmt.Println(x)

}