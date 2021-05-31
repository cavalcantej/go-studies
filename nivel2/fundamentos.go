package main 

import (
		"fmt"
		// "runtime"
)

// func main() {
// 	fmt.Println(runtime.GOOS)
// 	fmt.Println(runtime.GOARCH)

// 	s := "Hello, playground"

// 	fmt.Printf("%v\n%T", s, s)

// 	sb := []byte(s)


// 	for _, v := range sb {
// 		fmt.Printf("%v - %T %#U - %#x\n", v, v, v, v)
// 	}
	
// }


const(
	_ = iota
	KB = 1 << (iota * 10)
	MB
	GB
	TB
)

func main() {
	fmt.Println("binary \t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)

	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)

	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

	fmt.Printf("%b\t", TB)
	fmt.Printf("%d", TB)

}