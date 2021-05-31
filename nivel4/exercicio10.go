package main

import "fmt"

func main() {
	nomes := map[string][]string{
		"joao_zinh": []string{
			"crossfit", "something",
	},
		"Maria_can": []string{
			"vegana", "crossfit", "something",
		},
		"Josi_ildo":[]string{
			"crossfit", "something", "crossfit", "something",
		},
	}

	nomes["novo_nome"] = []string{
		"teste1", "teste2", 
	}

	for key, value := range nomes{
		fmt.Println(key)
		for k, v := range value{
			fmt.Printf("\t %v - %v\n",k, v)
		}
	}

	delete(nomes, "Maria_can")

	fmt.Println("///////////////////////////////////////////////")
	for key, value := range nomes{
		fmt.Println(key)
		for k, v := range value{
			fmt.Printf("\t %v - %v\n",k, v)
		}
	}
}