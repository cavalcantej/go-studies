package main 

import "fmt"

func main() {
	// Clojure é cercar ou capturar um scope para que possamos
	// utilizá-lo em outro contexto
	// exemplo
	//func i() func() int {x := 0; return func() int { x++; return x } }
	// Quando fizermos a := i() teremos um scope, um valor para x
	// Quando fizermos b := i() teremos outro scope, e x terá um valor independente do x acima.
	// Clojure nos permitem salvar dados entre function calls e ao mesmo tempo isolar estes dados do resto do código
	// função subjacente
	a := i()
	b := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())



}
func i() func() int {
	x := 0 

	return func() int { 
		x++
		return x 
	} 
}
