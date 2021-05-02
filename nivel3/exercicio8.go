package main 

import "fmt"


func main() {
	x := "testE"
	switch {
	case x == "teste" :
		fmt.Println("algo")
	case x == "test!":
		fmt.Println("algo2")
	case x == "testE":
		fmt.Println("algo3")
	}
}