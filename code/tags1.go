package main

import (
	"fmt"
	"reflect"
)

type Persona struct {
	Nombre string `validar:"largo=5"`
}

func main() {

	p := Persona{Nombre: "Test"}

	field, _ := reflect.TypeOf(p).FieldByName("Nombre")
	tag := field.Tag.Get("validar")

	fmt.Printf("el valor del tag es: '%v'", tag)
}
