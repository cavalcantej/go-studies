package main

import "fmt"

func main() {
	estados := make([]string, 27, 27)

	estados = []string{"Acre", "Alagoas", "Amapá"}

	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}

	fmt.Println(len(estados), cap(estados))
}