package main

import (
	"fmt"
	"reflect"
)

type Persona struct {
	Nombre string
}

func main() {
	p := new(Persona) // tipo *Persona
	var per Persona   // tipo  Persona
	fmt.Println("p: ", reflect.TypeOf(p), "per: ", reflect.TypeOf(per))
}
