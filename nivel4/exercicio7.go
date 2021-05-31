package main

import "fmt"

func main() {
	slice:= [][]string{
		[]string{
			"Joao", 
			"Feitosa", 
			"Cantar no chuveiro",
		},
		[]string{"Maria", "Josi", "Crossfit"},
		[]string{"Hulk", "Silv", "Tocar Guitarra"},
	}

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++{
			// fmt.Println(slice[i][j])
		}
	}

	for i, v := range slice{
		fmt.Printf("Slice %v\n", i)
		for _, item := range v{
			fmt.Println("\t", item)
		}
	}
}