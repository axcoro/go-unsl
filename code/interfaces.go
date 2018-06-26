package main

import "fmt"

type Animal interface {
	Decir()
}

func hacerHablar(a Animal) {
	a.Decir()
}

type Pato struct{}
func (p Pato) Decir() {
	fmt.Printf("cuack cuack!")
}

func main() {
	p := Pato{}
	hacerHablar(p)
}
