package main

import "fmt"

func soma(x ...int) int {
	total := 0
	for _, v := range x{
		total += v
	}
	return total
}

func somaslice(x []int) int {
	total := 0
	for _, v := range x{
		total += v
	}
	return total
}

func somatudo(x ...interface{}) int {
	s := make([]int, 0, len(x))

	for _, v := range x {
		item_type := fmt.Sprintf("%T", v)

		if item_type == "int"{ 
			s = append(s, v.(int))
		} else if item_type == "[]int" {
			s = append(s,v.([]int)...)
		}
	}
	
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}


func main(){
	fmt.Println(soma(1,2,3,4,5))
	slice := []int{1,2,3,4,5}
	
	fmt.Println(somaslice(slice))

	fmt.Println(somatudo(1,2,3,4, []int{10,43,43,43}))


}