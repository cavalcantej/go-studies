package main

import "fmt"

type pessoa struct{
	nome string
	idade int
}
//método
func (p pessoa) oibomdia(){
	fmt.Println(p.nome, "diz bom dia!")
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func main() {
	s1 := []int{5, 2, 4, 7, 8, 9, 11}
	total := soma(s1...)
	fmt.Println(total)

// defer -> deixa pra ultima hora, se tiver return, acontece antes do return
	fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")	

// métodos
// método é uma função anexada a um tipo
	mauricio := pessoa{"Mauricio", 30}
	mauricio.oibomdia()

}