package main

import "fmt"


func multiplica(x ...int) int{
	n := 1
	for _, v := range x{
		n *= v
	}
	return n
}

func somenteImpares(f func(x ...int) int, y ...int) int{
	var slice []int

	for _, v := range y {
		if v % 2 == 1 {
			slice = append(slice, v)
		}
	}

	total := f(slice...)
	return total
}

func main(){
	// o que é um callback?
	// callback é uma função que retorna outra função
	// a função recebe uma função como argumento, e quaisquer outro argumento
	// a função principal é executada, e por ultima a função passada como argumento é executada

	t := somenteImpares(multiplica, []int{51,52,53,54,55,67,68,60,59,71}...)
	fmt.Println(t)
}