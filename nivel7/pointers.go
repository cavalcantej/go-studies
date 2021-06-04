package main

import "fmt"

func main() {
	x := 0
	
	// pega o endereço da memória
	y := &x

	
	// dereference -> tenho um endereço e quero saber o que tem lá dentro daquele endereço
	// conteudo do endereço da variavel x
	fmt.Println("Valor da variável x:", *&x)

	// printa o endereço da variável x
	fmt.Println("Valor da variável y, endereço de memória de x: ",y)

	// o endereço de y é diferente de x, porém o valor de y é o endereço de x
	fmt.Println("Endereço da variável y, != endereço de x: ",&y)
	
	// valor contido no endereço guardado em y
	fmt.Println("Valor de x, *y:",*y)

	// demonstração dos tipos x e y
	fmt.Printf("Tipo de x: %T\nTipo de y: %T\n", x, y)

	// modificar o valor que está no endereço da memória
	*y++

	fmt.Println("Após *y++, valor de x: ", x)

	// func (x *int)
	// *x++

}

