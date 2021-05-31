package main 

import "fmt"


func main() {
	// x := 3
	// fmt.Println(half(x))

	// fmt.Println(greatest(13,14,15,16,23,818,99,103,114))
	// nextOdd := makeOddGenerator()
	
	// for i := 0; i < 10; i++{
	// 	fmt.Println(nextOdd())
	// }

	// fmt.Println(fib(10))

	// defer is used when resources need to be freed in some way, like, when a function depends on other to occur
	// in that way, we can combine it's functionality with recover and panic
	// so recover is called after the panic function occurs
	// using defer we can ensure that the recover function is called when panic occurs
	// panic stops immediately the execution of the function
	// defer func() {
	// 	str := recover()
	// 	fmt.Println("something went wrong: ", str)
	// }()
	// panic("PANIC")
	
	// newPointer := new(int)
	// pointerTo(&x)
	// fmt.Println(*newPointer)

	x := 3
	y := 5

	fmt.Println(swap(&x, &y))

}


func half(x int) (int, bool){
	if x % 2 == 0 {
		return 1, true
	} else {
		return 0, false
	}
}


func greatest(args ...int) int{
	n := 0
	for _, v := range args {
		if n < v{
			n = v
		}
	}
	return n
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return 
	}
}

func fib(x int) int{
	if x == 0 {
		return 0
	}

	return x + fib(x - 1)
}

func pointerTo(xPointer *int){
	*xPointer = 0
}

func swap(x *int, y *int) (int, int){
	holder := *y
	*y = *x
	*x = holder
	return *x, *y
}
