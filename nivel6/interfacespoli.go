package main

import "fmt"

type pessoa struct{
	nome		string
	sobrenome 	string
	idade 		int
}

type dentista struct {
	pessoa 
	dentesarrancados int
	salário int
}

type arquiteto struct {
	pessoa
	tipodeconstrução string
	tamanhodaloucura string
}

func (x dentista) oibomdia(){
	fmt.Println("Meu nome é ", x.nome, "e eu já arranquei", x.dentesarrancados, "ouve só: Bom dia!")
}

func (x arquiteto) oibomdia(){
	fmt.Println("Meu nome é ", x.nome, "e ouve só: Bom dia!")
}

type gente interface{
	oibomdia()
}

func serhumano(g gente){
	g.oibomdia()
}

func main(){
	mrdente := dentista{
		pessoa: pessoa{
			nome: "Mister Dente",
			sobrenome: "da Silva",
			idade: 50,
		},
		dentesarrancados: 99102,
		salário: 2123213123,
	}

	mrpredio := arquiteto{
		pessoa: pessoa{
			nome: "Mister Construção",
			sobrenome: " de Tijolos",
			idade: 70,
		},
		tipodeconstrução: "Hospícios",
		tamanhodaloucura: "Demais",
	}
}