package main 

import "fmt"

// criar tipos : quadrado, circulo

type quadrado struct {
	lado float64
}

type circulo struct{
	raio float64
}

// criar método "area" que retorne a area de cada forma geométrica

func (q quadrado) area() float64{
	return q.lado * q.lado
}

func (c circulo) area() float64{
	return 2 * 3.14 * c.raio
}

// criar tipo "figura" que defina como interface qualquer tipo que tiver o método área
type figura interface{
	area() float64
}

// crie uma função "info" que receba um tipo "figura" e retorna a área da figura
func info(f figura){
	fmt.Println(f.area())
}

func main() {
	// crie um valor de tipo "quadrado"
	q := quadrado{lado:10}
	// crie um valor de tipo "círculo"
	c := circulo{raio:12}
	// use a função "info" para demonstrar a área do "quadrado"
	info(q)	
	// use a função "info" para demonstrar a área do "círculo"
	info(c)

}