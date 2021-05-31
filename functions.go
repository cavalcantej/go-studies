package main

import "fmt"


func average(xs []float64) float64{ // this are the functions signatures
	// parameters & return type
	// panic("Not Implemented")

	total := 0.0

	for _, value := range xs {
		total += value
	}

	return total / float64(len(xs))
}

// calling and callee functions
// functions are built up in a call stack


// multiple values can be returned 
func f2() (int, int) {
	return 5, 6
}

// variadic functions
// those are functions where the parameter can receive many arguments
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}


func factorial(x uint) uint{
	if x == 0 {
		return 1 
	}
	return x * factorial(x - 1)
}


// // defer, panic, and recover
// f, _ := os.Open(filename)
// defer f.Close()

// defer func() {
// 	str := recover()
// 	fmt.Println(str)
// }()
// panic("PANIC")


func main() {
	defer func() {
		str := recover()
		fmt.Println("something went wrong: ", str)
	}()
	panic("PANIC")
	
	fmt.Println(factorial(5))
    nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) 

	// runtime functions 
	// add2 := func(x, y int) int {
    //     return x + y
    // }
	// fmt.Println(add2(1,2))

	// add3 := func(args ...int) int {
	// 	total := 0
	// 	for _, v := range args {
	// 		total+= v
	// 	}
	// 	return total
	// }
	// fmt.Println(add3(1,2,3,4,5,6))


    // fmt.Println(add(1,1,2,3,4,5,6,7,8,9,1))

	// xs := []int{2, 34, 56}
	// fmt.Println(add(xs...))

    
}