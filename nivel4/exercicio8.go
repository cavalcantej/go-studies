package main

import "fmt"

func main() {
	nomes := map[string]string{
		"joao_zinho": "crossfit",
		"Maria_carmen": "vegana",
		"Josi_ildo": "pedreiro", 
	}

	fmt.Println(nomes)

	for key, value := range nomes{
		fmt.Println(key, value)
	}
}