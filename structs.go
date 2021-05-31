package main

import ("fmt"; "math")

type Circle struct {
	x, y, r float64

}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
    area() float64
}

type MultiShape struct {
    shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
func (c *Circle) area() float64{
	return math.Pi * c.r*c.r
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}
func main() {

	multiShape := MultiShape{
		shapes: []Shape{
			Circle{0, 0, 5},
			Rectangle{0, 0, 10, 10},
		},
	}

	fmt.Println(multiShape.area())


	// // instances
	// c1 := new(Circle)

	// // attribution
	// c := Circle{x: 0, y: 0, r: 5}
	// c1 := &Circle{0,0,5}
	
	// // accessing the values
	// fmt.Println(c.x, c.y, c.r)
	// c.x = 10
	// c.y = 5
	// fmt.Println(c.x, c.y, c.r)

	// fmt.Println(c1.x, c1.y, c1.r)

	// c := Circle{0,0,5}

	// // area := circleArea(&c)

	// fmt.Println(c.area())

	// r := Rectangle{0, 0, 10, 10}
	// fmt.Println(r.area())

	
	

}
