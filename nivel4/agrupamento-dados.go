package main

import "fmt"

func main(){
// ARRAYS
	// var x [5]int
	// var y [6]int
	// // atribuindo valores
	// x[0] = 1
	// x[1] = 10

	// // funções para array len()
	// fmt.Println(len(x))


// SLICES
	// slice := []int{1, 5, 7, 9}

	// slice = append(slice, 7)

	// fmt.Println(len(slice))

	// for i, v := range slice{
	// 	fmt.Println(i, v)
	// }

// SLICING SLICES

	// sabores := []string{"3123213", "213123", "3123213", "1321312", "3123123 queijos"}
	
	// // fatia := sabores[:]

	// semabacaxi := append(sabores[0:3], sabores[4:]...)

	// fmt.Println(semabacaxi)

// INDEXING SLICES
	// sabores2 := []string{"pepperoni", "mozzarela", "margueritta", "abacaxi", "quatro queijos"}


	// sabores2 = append(sabores2, sabores...)

	// fmt.Println(sabores2)
// MAKE FUNCTION
	// 	slice := make([]int, 5, 10)

	// 	slice[0], slice[1], slice[2], slice[3], slice[4] = 0,1,2,3,4 

	// 	fmt.Println(slice, len(slice), cap(slice))

// MAPS
	amigos := map[string]int {
		"joaozinho" : 0,
		"maria" : 202020,
	}

	// fmt.Println(amigos["joaozinho"])

	if teste, ok := amigos["a"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(teste)
	}
 
	for key, value := range amigos{
		fmt.Println(key, value)
	}

	delete(amigos, "joaozinho")
	
	fmt.Println(amigos)

}