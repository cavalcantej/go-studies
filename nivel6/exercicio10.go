package main

import "fmt"

func cloj() func() int{
	x := 10
	y := 4

	return func() int {
		x = x * y
		return x
	}

}

func main(){
	a := cloj()
	b := cloj()


	fmt.Println("essa é a ", a())

	fmt.Println("essa é b ", b())	
	fmt.Println("essa é b ", b())	
	fmt.Println("essa é b ", b())	
	fmt.Println("essa é b ", b())	

	
	fmt.Println("essa é a ", a())

	fmt.Println("essa é b ", b())
	fmt.Println("essa é a ", a())

	fmt.Println("essa é b ", b())
	fmt.Println("essa é a ", a())

	fmt.Println("essa é b ", b())
	fmt.Println("essa é a ", a())

	fmt.Println("essa é b ", b())

}

