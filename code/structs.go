package main

import "fmt"

type Persona struct {
	ID       int
	Nombre   string
	Apellido string
	Email    string
}

func (p Persona) Hablar() {
	fmt.Printf("Hola, soy %s", p.Nombre)
}

func main() {

	p := Persona{Nombre: "Test", Apellido: "User"}

	fmt.Printf("%s\n", p.Nombre)
	fmt.Println("Apellido:", p.Apellido)

	p.Hablar()
}
