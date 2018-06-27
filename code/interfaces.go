package main

import "fmt"

type Animal interface {
	HacerRuido()
}

func hacerRuido(a Animal) {
	a.HacerRuido()
}

type Pato struct{}
func (p Pato) HacerRuido() {
	fmt.Printf("cuack cuack!")
}

func main() {
	p := Pato{}
	hacerRuido(p)
}
