package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is ", p.Name)
}

func main() {
	android18 := new(Android)
	android18.Name = "neo"
	android18.Talk()

}