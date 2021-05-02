package main

import "fmt"

func main() {
	

	for x := 1; x <= 10; x++ {
		if x % 2 == 0 {
			fmt.Println(x, "even")
		} else {
			fmt.Println(x, "odd")
		}

		switch x {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			case 4: fmt.Println("Four")
			case 5: fmt.Println("Five")
			default: fmt.Println("Unknown Number")
		}

	}



}