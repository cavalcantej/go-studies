package main 

import "fmt"


func main() {
	x := 11

	if x < 10 {
		fmt.Println("Dentro do if")
	} else if x == 10 {
		fmt.Println("Dentro do else if")
	} else {
		fmt.Println("Dentro do else")
	}

}